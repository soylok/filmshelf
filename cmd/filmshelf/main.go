package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/media"
	"github.com/soylok/filmshelf/internal/core/repository/postgres/pool"
	"github.com/soylok/filmshelf/internal/core/transport/middleware"
	"github.com/soylok/filmshelf/internal/core/transport/server"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/filesystem/posters"
	countryhandler "github.com/soylok/filmshelf/internal/features/catalog/adapters/http/countries"
	directorhandler "github.com/soylok/filmshelf/internal/features/catalog/adapters/http/directors"
	genrehandler "github.com/soylok/filmshelf/internal/features/catalog/adapters/http/genres"
	moviehandler "github.com/soylok/filmshelf/internal/features/catalog/adapters/http/movies"
	countryrepo "github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/countries"
	directorrepo "github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/directors"
	genrerepo "github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/genres"
	movierepo "github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/movies"
	countrysvc "github.com/soylok/filmshelf/internal/features/catalog/service/countries"
	directorsvc "github.com/soylok/filmshelf/internal/features/catalog/service/directors"
	genresvc "github.com/soylok/filmshelf/internal/features/catalog/service/genres"
	moviesvc "github.com/soylok/filmshelf/internal/features/catalog/service/movies"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger, err := logger.NewLogger(logger.NewConfigMust())
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	mediaCfg, err := media.NewConfig()
	if err != nil {
		panic(err)
	}
	if err := media.EnsurePostersDir(mediaCfg.PostersDir); err != nil {
		panic(err)
	}

	poolCfg := pool.NewConfigMust()
	pool, err := pool.NewPool(ctx, poolCfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	countryRepo := countryrepo.NewCountryRepository(pool, poolCfg.OpTimeout)
	countrySvc := countrysvc.NewCountryService(countryRepo)
	countryHandler := countryhandler.NewCountryHandler(countrySvc)

	directorRepo := directorrepo.NewDirectorRepository(pool, poolCfg.OpTimeout)
	directorSvc := directorsvc.NewDirectorService(directorRepo)
	directorHandler := directorhandler.NewDirectorHandler(directorSvc)

	genreRepo := genrerepo.NewGenreRepository(pool, poolCfg.OpTimeout)
	genreSvc := genresvc.NewGenreService(genreRepo)
	genreHandler := genrehandler.NewGenreHandler(genreSvc)

	movieRepo := movierepo.NewMovieRepository(pool, poolCfg.OpTimeout)
	posterStorage := posters.NewPosterStorage(mediaCfg.PostersDir)
	movieSvc := moviesvc.NewMovieService(movieRepo, posterStorage)
	movieHandler := moviehandler.NewMovieHandler(movieSvc, mediaCfg.PostersDir)

	vRouter := server.NewRouter(server.Version1)
	vRouter.RegisterRoutes(countryHandler.Routes())
	vRouter.RegisterRoutes(directorHandler.Routes())
	vRouter.RegisterRoutes(genreHandler.Routes())
	vRouter.RegisterRoutes(movieHandler.Routes())

	HTTPServer := server.NewHTTPServer(
		server.NewConfigMust(),
		logger,
		[]middleware.Middleware{
			middleware.InjectLogger(logger),
			middleware.RequestID(),
			middleware.Trace(),
			middleware.Recovery(),
		},
	)

	HTTPServer.RegisterRouters(vRouter)
	if err := HTTPServer.Run(ctx); err != nil {
		panic(err)
	}
}

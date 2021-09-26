// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ccremer/greposync/application"
	"github.com/ccremer/greposync/application/initialize"
	"github.com/ccremer/greposync/application/instrumentation"
	"github.com/ccremer/greposync/application/labels"
	"github.com/ccremer/greposync/application/update"
	"github.com/ccremer/greposync/cfg"
	"github.com/ccremer/greposync/domain"
	"github.com/ccremer/greposync/infrastructure/githosting"
	"github.com/ccremer/greposync/infrastructure/githosting/github"
	"github.com/ccremer/greposync/infrastructure/repositorystore"
	"github.com/ccremer/greposync/infrastructure/templateengine"
	"github.com/ccremer/greposync/infrastructure/templateengine/gotemplate"
	"github.com/ccremer/greposync/infrastructure/ui"
	"github.com/ccremer/greposync/infrastructure/valuestore"
)

// Injectors from wire.go:

func initInjector() *injector {
	versionInfo := _wireVersionInfoValue
	configuration := cfg.NewDefaultConfig()
	coloredConsole := ui.NewColoredConsole()
	consoleSink := ui.NewConsoleSink(coloredConsole)
	consoleLoggerFactory := ui.NewConsoleLoggerFactory(consoleSink)
	repositoryStoreInstrumentation := repositorystore.NewRepositoryStoreInstrumentation(consoleLoggerFactory)
	repositoryStore := repositorystore.NewRepositoryStore(repositoryStoreInstrumentation)
	ghRemote := github.NewRemote(consoleLoggerFactory)
	providerMap := newGitProviders(ghRemote)
	labelStore := githosting.NewLabelStore(providerMap)
	appService := labels.NewConfigurator(repositoryStore, labelStore, configuration, consoleLoggerFactory)
	commonBatchInstrumentation := instrumentation.NewUpdateInstrumentation(coloredConsole, consoleLoggerFactory)
	command := labels.NewCommand(configuration, appService, commonBatchInstrumentation)
	goTemplateEngine := gotemplate.NewEngine()
	goTemplateStore := gotemplate.NewTemplateStore()
	valueStoreInstrumentation := valuestore.NewValueStoreInstrumentation(consoleLoggerFactory)
	koanfValueStore := valuestore.NewValueStore(valueStoreInstrumentation)
	pullRequestStore := githosting.NewPullRequestStore(providerMap)
	renderServiceInstrumentation := templateengine.NewRenderServiceInstrumentation(consoleLoggerFactory)
	renderService := domain.NewRenderService(renderServiceInstrumentation)
	consoleDiffPrinter := ui.NewConsoleDiffPrinter()
	updateAppService := update.NewConfigurator(goTemplateEngine, repositoryStore, goTemplateStore, koanfValueStore, pullRequestStore, renderService, consoleDiffPrinter, configuration, coloredConsole)
	updateCommand := update.NewCommand(configuration, updateAppService, consoleLoggerFactory, commonBatchInstrumentation)
	initializeCommand := initialize.NewCommand(configuration, consoleLoggerFactory)
	app := application.NewApp(versionInfo, configuration, command, updateCommand, initializeCommand, consoleLoggerFactory)
	mainInjector := NewInjector(app)
	return mainInjector
}

var (
	_wireVersionInfoValue = application.VersionInfo{Version: version, Commit: commit, Date: date}
)

// wire.go:

type injector struct {
	app *application.App
}

func NewInjector(
	app *application.App,
) *injector {
	i := &injector{
		app: app,
	}
	return i
}

func (i *injector) RunApp() {
	i.app.Run()
}

func newGitProviders(ghRemote *github.GhRemote) githosting.ProviderMap {
	return githosting.ProviderMap{github.ProviderKey: ghRemote}
}

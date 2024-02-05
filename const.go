package cidsdk

// see: https://github.com/cidverse/repoanalyzer/blob/main/analyzerapi/types.go

type ProjectLanguage string

const (
	LanguageGolang     ProjectLanguage = "go"
	LanguageJava       ProjectLanguage = "java"
	LanguageKotlin     ProjectLanguage = "kotlin"
	LanguageJavascript ProjectLanguage = "javascript"
	LanguageTypescript ProjectLanguage = "typescript"
	LanguagePython     ProjectLanguage = "python"
	LanguagePHP        ProjectLanguage = "php"
	LanguageRust       ProjectLanguage = "rust"
)

type ProjectBuildSystem string

const (
	BuildSystemGradle          ProjectBuildSystem = "gradle"
	BuildSystemMaven           ProjectBuildSystem = "maven"
	BuildSystemGoMod           ProjectBuildSystem = "gomod"
	BuildSystemNpm             ProjectBuildSystem = "npm"
	BuildSystemHugo            ProjectBuildSystem = "hugo"
	BuildSystemHelm            ProjectBuildSystem = "helm"
	BuildSystemContainer       ProjectBuildSystem = "container"
	BuildSystemRequirementsTXT ProjectBuildSystem = "python-requirements.txt"
	BuildSystemPipfile         ProjectBuildSystem = "pipfile"
	BuildSystemSetupPy         ProjectBuildSystem = "setup.py"
	BuildSystemMkdocs          ProjectBuildSystem = "mkdocs"
	BuildSystemComposer        ProjectBuildSystem = "composer"
	BuildSystemDotNet          ProjectBuildSystem = "dotnet"
	BuildSystemCargo           ProjectBuildSystem = "cargo"
)

type ProjectBuildSystemSyntax string

const (
	BuildSystemSyntaxDefault ProjectBuildSystemSyntax = "default"
	GradleGroovyDSL          ProjectBuildSystemSyntax = "groovy"
	GradleKotlinDSL          ProjectBuildSystemSyntax = "kotlin"
	ContainerFile            ProjectBuildSystemSyntax = "containerfile"
	ContainerBuildahScript   ProjectBuildSystemSyntax = "buildah-script"
	MkdocsTechdocs           ProjectBuildSystemSyntax = "mkdocs-techdocs"
)

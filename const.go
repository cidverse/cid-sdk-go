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
	LanguageNix        ProjectLanguage = "nix"
	LanguageOpenAPI    ProjectLanguage = "openapi"
	LanguageAsyncAPI   ProjectLanguage = "asyncapi"
)

type ProjectBuildSystem string

const (
	BuildSystemDefault         ProjectBuildSystem = "default"
	BuildSystemGradle          ProjectBuildSystem = "gradle"
	BuildSystemMaven           ProjectBuildSystem = "maven"
	BuildSystemGoMod           ProjectBuildSystem = "gomod"
	BuildSystemNpm             ProjectBuildSystem = "npm"
	BuildSystemHugo            ProjectBuildSystem = "hugo"
	BuildSystemHelm            ProjectBuildSystem = "helm"
	BuildSystemHelmfile        ProjectBuildSystem = "helmfile"
	BuildSystemContainer       ProjectBuildSystem = "container"
	BuildSystemRequirementsTXT ProjectBuildSystem = "python-requirements.txt"
	BuildSystemPipfile         ProjectBuildSystem = "pipfile"
	BuildSystemSetupPy         ProjectBuildSystem = "setup.py"
	BuildSystemPoetry          ProjectBuildSystem = "poetry"
	BuildSystemMkdocs          ProjectBuildSystem = "mkdocs"
	BuildSystemComposer        ProjectBuildSystem = "composer"
	BuildSystemDotNet          ProjectBuildSystem = "dotnet"
	BuildSystemCargo           ProjectBuildSystem = "cargo"
	BuildSystemNix             ProjectBuildSystem = "nix"
	BuildSystemAnsible         ProjectBuildSystem = "ansible"
)

type ProjectBuildSystemSyntax string

const (
	BuildSystemSyntaxDefault                ProjectBuildSystemSyntax = "default"
	BuildSystemSyntaxGradleGroovyDSL        ProjectBuildSystemSyntax = "groovy"
	BuildSystemSyntaxGradleKotlinDSL        ProjectBuildSystemSyntax = "kotlin"
	BuildSystemSyntaxDotNetSLN              ProjectBuildSystemSyntax = "sln"
	BuildSystemSyntaxDotNetSLNX             ProjectBuildSystemSyntax = "slnx"
	BuildSystemSyntaxDotNetCSProj           ProjectBuildSystemSyntax = "csproj"
	BuildSystemSyntaxContainerFile          ProjectBuildSystemSyntax = "containerfile"
	BuildSystemSyntaxContainerBuildahScript ProjectBuildSystemSyntax = "buildah-script"
	BuildSystemSyntaxMkdocsTechdocs         ProjectBuildSystemSyntax = "mkdocs-techdocs"
	BuildSystemSyntaxNixFlake               ProjectBuildSystemSyntax = "flake"
)

type SpecificationType string

const (
	SpecificationTypeOpenAPI  SpecificationType = "openapi"
	SpecificationTypeAsyncAPI SpecificationType = "asyncapi"
	SpecificationTypeRenovate SpecificationType = "renovate"
)

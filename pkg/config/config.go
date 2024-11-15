package config

type EventConfiguration struct {
	ManifestId       int64
	ScenarioId       int64
	GroupStageConfig GroupStageConfig
	PlayOffConfig    PlayOffConfig
}

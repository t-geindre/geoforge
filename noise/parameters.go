package noise

import "geoforge/preset"

const (
	SeedMin = 0
	SeeMax  = 2147483647
)

const (
	// Noise manager
	ParamIsRendered   preset.ParamId = 100
	ParamName         preset.ParamId = 101
	ParamActionAdd    preset.ParamId = 102
	ParamActionRemove preset.ParamId = 103
	ParamType         preset.ParamId = 104

	// Noise parameters
	ParamSeed            preset.ParamId = 110
	ParamSubType         preset.ParamId = 111
	ParamScale           preset.ParamId = 112
	ParamSetFract        preset.ParamId = 113
	ParamFractType       preset.ParamId = 114
	ParamFractOctaves    preset.ParamId = 115
	ParamFractLacunarity preset.ParamId = 116
	ParamFractGain       preset.ParamId = 117
	ParamFractWeighting  preset.ParamId = 118
	ParamFractPPStrength preset.ParamId = 123
	ParamSetWarp         preset.ParamId = 119
	ParamWarpType        preset.ParamId = 120
	ParamWarpAmp         preset.ParamId = 121
	ParamWarpFreq        preset.ParamId = 122
)

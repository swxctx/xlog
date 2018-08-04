package xlog

// Level
type Level int

const (
	// Disable non print
	DisableLevel = iota
	// Fatal os.exit()
	FatalLevel
	// Error print only errors.
	ErrorLevel
	// Warn print errors and warnings.
	WarnLevel
	// Info print errors, warnings and infos.
	InfoLevel
	// Debug will print on any level
	DebugLevel
)

// Levels info
var Levels = map[Level]*LevelMeta{
	DisableLevel: {
		Name:             "disable",
		AlternativeNames: []string{"disabled"},
		RawText:          "",
		ColorfulText:     "",
	},
	FatalLevel: {
		Name:         "fatal",
		RawText:      "[FTAL]",
		ColorfulText: LightGreen("[FTAL]"),
	},
	ErrorLevel: {
		Name:         "error",
		RawText:      "[ERRO]",
		ColorfulText: Red("[ERRO]"),
	},
	WarnLevel: {
		Name:             "warn",
		AlternativeNames: []string{"warning"},
		RawText:          "[WARN]",
		ColorfulText:     Purple("[WARN]"),
	},
	InfoLevel: {
		Name:         "info",
		RawText:      "[INFO]",
		ColorfulText: Green("[INFO]"),
	},
	DebugLevel: {
		Name:         "debug",
		RawText:      "[DBUG]",
		ColorfulText: Yellow("[DBUG]"),
	},
}

// getLevelByName
func getLevelByName(levelName string) Level {
	for level, meta := range Levels {
		if meta.Name == levelName {
			return level
		}

		for _, altName := range meta.AlternativeNames {
			if altName == levelName {
				return level
			}
		}
	}
	return DisableLevel
}

// LevelMeta
type LevelMeta struct {
	Name             string
	AlternativeNames []string
	RawText          string
	ColorfulText     string
}

func (m *LevelMeta) Text(enableColor bool) string {
	if enableColor {
		return m.ColorfulText
	}
	return m.RawText
}

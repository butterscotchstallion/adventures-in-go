package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron/v2"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const appName string = "Disk Space Monitor"

func CheckLowSpaceAndNotify() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	fmt.Println(appName)

	// Get low space devices
	diskPartitions, _ := disk.Partitions(true)
	lowSpaceDrives := GetLowDiskSpaceDrives(diskPartitions, 90.0)
	messages := make([]string, 0)
	for _, statInfo := range lowSpaceDrives {
		message := fmt.Sprintf("%v has low disk space (%v%% usage)\n", statInfo.MountPoint, statInfo.DiskUsagePercent)
		sugar.Info(message)
		messages = append(messages, message)
	}

	// Show notifications, if any low space drives
	if len(messages) > 0 {
		sugar.Debugf("Showing %v notifications", len(messages))
		for _, message := range messages {
			ShowNotification("Low disk space", message)
		}
	}
}

func getIconBytes(logger *zap.SugaredLogger) ([]byte, error) {
	iconFilePath := "disk_space_monitor_icon.ico"
	bytes, err := os.ReadFile(iconFilePath)
	if err != nil {
		logger.Fatalf("Can't load icon file: %v", err)
	}
	return bytes, nil
}

func onReady(logger *zap.SugaredLogger) {
	iconBytes, _ := getIconBytes(logger)
	systray.SetIcon(iconBytes)
	systray.SetTitle(appName)
	systray.SetTooltip("Monitoring disk space...")
	mQuit := systray.AddMenuItem("Quit", "Quit and stop monitoring")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	os.Exit(0)
}

func setSystemTrayIcon(logger *zap.SugaredLogger) {
	var onReadyCallback = func() {
		onReady(logger)
	}
	systray.Run(onReadyCallback, onExit)
}

func scheduleSpaceCheck(logger *zap.SugaredLogger) {
	scheduler, _ := gocron.NewScheduler()
	ScheduleSpaceCheck(scheduler, logger)
}

func createLogger() (*zap.SugaredLogger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}

	// Set up lumberjack as a logger:
	logger := &lumberjack.Logger{
		Filename:   "./logs/dsm.log", // Or any other path
		MaxSize:    500,              // MB; after this size, a new log file is created
		MaxBackups: 3,                // Number of backups to keep
		MaxAge:     28,               // Days
		Compress:   true,             // Compress the backups using gzip
	}

	writeSyncer := zapcore.AddSync(logger)

	// Set up zap logger configuration:
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // Using JSON encoder, but you can choose another
		writeSyncer,
		zapcore.DebugLevel,
	)

	loggerZap := zap.New(core)
	sugar := loggerZap.Sugar()
	defer sugar.Sync()

	return sugar, nil
}

func main() {
	sugar, err := createLogger()
	if err != nil {
		sugar.Fatal(err)
	}

	CheckLowSpaceAndNotify()
	go scheduleSpaceCheck(sugar)
	setSystemTrayIcon(sugar)
}

package main

var SessionGetBase = JsonMap{
	"alt-speed-down": 50,
	"alt-speed-enabled": false,
	"alt-speed-time-begin": 540,
	"alt-speed-time-day": 127,
	"alt-speed-time-enabled": false,
	"alt-speed-time-end": 1020,
	"alt-speed-up": 50,
	"blocklist-enabled": false,
	"blocklist-size": 393006,
	"blocklist-url": "http://www.example.com/blocklist",
	"cache-size-mb": 4,
	"config-dir": "/var/lib/transmission-daemon",
	"dht-enabled": true,
	"download-dir": "/var/lib/transmission-daemon/downloads",
	"download-dir-free-space": float64(100 * (1 << 30)), // 100 GB
	"download-queue-enabled": true,
	"download-queue-size": 5,
	"encryption": "preferred",
	"idle-seeding-limit": 30,
	"idle-seeding-limit-enabled": false,
	"incomplete-dir": "/var/lib/transmission-daemon/downloads",
	"incomplete-dir-enabled": false,
	"lpd-enabled": false,
	"peer-limit-global": 200,
	"peer-limit-per-torrent": 50,
	"peer-port": 44444,
	"peer-port-random-on-start": false,
	"pex-enabled": true,
	"port-forwarding-enabled": true,
	"queue-stalled-enabled": true,
	"queue-stalled-minutes": 30,
	"rename-partial-files": true,
	"rpc-version": 15,
	"rpc-version-minimum": 1,
	"script-torrent-done-enabled": false,
	"script-torrent-done-filename": "",
	"seed-queue-enabled": false,
	"seed-queue-size": 10,
	"seedRatioLimit": 2,
	"seedRatioLimited": false,
	"speed-limit-down": 100,
	"speed-limit-down-enabled": false,
	"speed-limit-up": 100,
	"speed-limit-up-enabled": false,
	"start-added-torrents": true,
	"trash-original-torrent-files": false,
	"units":map[string]interface{}{
		"memory-bytes":1024,
		"memory-units":[]string{
			"KiB",
			"MiB",
			"GiB",
			"TiB",
		},
		"size-bytes":1000,
		"size-units":[]string{
			"kB",
			"MB",
			"GB",
			"TB",
		},
		"speed-bytes":1000,
		"speed-units":[]string{
			"kB/s",
			"MB/s",
			"GB/s",
			"TB/s",
		},
	},
	"utp-enabled": false,
	"version": "2.84 (14307)",
}

var TorrentGetBase = JsonMap{
	"errorString": "",
	"metadataPercentComplete": 0,
	"peersGettingFromUs": 0,
	"peersSendingToUs": 0,
	"pieces": "///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////w",
	"downloadedEver": 1 * (1 << 30), // 1 GB,
	"uploadedEver": 1 * (1 << 30), // 1 GB,
	"error": 0, // TR_STAT_OK
	"isFinished": false,
	"isStalled": false,
	"peersConnected": 0,
	"percentDone": 0,
	"queuePosition": 0,
	"seedRatioLimit": 2,
	"seedRatioMode": 0,
	"activityDate": 1443977197,
	"corruptEver": 0,
	"downloadLimited": false,
	"maxConnectedPeers": 50,
	"secondsDownloading": 500,
	"secondsSeeding": 80000,
	"uploadLimited": false,
	"isPrivate": false,
	"honorsSessionLimits": true,
	"webseedsSendingToUs": 0,
	"peer-limit": 50,
	"bandwidthPriority": 0,
	"seedIdleLimit": 10,
	"seedIdleMode": 0,
	// TODO
	"peers" : []string{},
}

var trackerStatsTemplate = JsonMap{
	"announceState": 0,
	"downloadCount": -1,
	"hasAnnounced": false,
	"hasScraped": false,
	"host": "http://example.com:80",
	"isBackup": false,
	"lastAnnouncePeerCount": 0,
	"lastAnnounceResult": "",
	"lastAnnounceStartTime": 0,
	"lastAnnounceSucceeded": false,
	"lastAnnounceTime": 0,
	"lastAnnounceTimedOut": false,
	"lastScrapeResult": "",
	"lastScrapeStartTime": 0,
	"lastScrapeSucceeded": false,
	"lastScrapeTime": 0,
	"lastScrapeTimedOut": 0,
	"leecherCount": -1,
	"nextAnnounceTime": 0,
	"nextScrapeTime": 0,
	"scrapeState": 2,
	"seederCount": -1,
}

var SessionStatsTemplate = JsonMap{
	"activeTorrentCount": 0,
	"cumulative-stats": map[string]int64{
		"downloadedBytes": 388802690736,
		"filesAdded": 5611,
		"secondsActive": 15681693897,
		"sessionCount": 57,
		"uploadedBytes": 1950265820985,
	},
	"current-stats": map[string]int64{
		"downloadedBytes": 9939147143,
		"filesAdded": 13,
		"secondsActive": 99633,
		"sessionCount": 1,
		"uploadedBytes": 26478335758,
	},
	"downloadSpeed": 0,
	"pausedTorrentCount": 127,
	"torrentCount": 127,
	"uploadSpeed": 0,
}
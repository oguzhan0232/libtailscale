// UYTE yaması: Tailscale log yüklemesini (log.tailscale.com) kapatır.
// Hasta verisi içeren bir ortamda hiçbir tanılama verisi üçüncü tarafa gitmemeli.
// Bu dosya Vendor/libtailscale/ içine kopyalanır (bkz. CLAUDE.md).
package main

import "os"

func init() {
	os.Setenv("TS_NO_LOGS_NO_SUPPORT", "true")
}

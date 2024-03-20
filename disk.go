package diskspace

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"net/url"
	"runtime"
	"strconv"
)

type DiskSpace struct {
	Telegram ITelegram
	Cfg      DiskConfig
}

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func FormatFileSize(s float64, base float64) string {
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

func (diskObj *DiskSpace) WarnDiskSpace() {
	runtimeOS := runtime.GOOS
	var param = url.Values{}

	diskStat, err := disk.Usage(diskObj.Cfg.DiskPath)
	if err != nil {
		param.Set("text", "cannot read path "+diskObj.Cfg.DiskPath+" "+err.Error())
		diskObj.Telegram.SendMessage(param)
		return
	}

	html := "<b>PERINGATAN!!</b>\n"

	if diskStat.UsedPercent > float64(diskObj.Cfg.MaxPercentage) {
		html = html + "------------------------------------------\n"
		html = html + "OS : <b>" + runtimeOS + "</b>\n"
		html = html + "Path : <b>" + diskStat.Path + "</b>\n"
		html = html + "Total disk space: " + FormatFileSize(float64(diskStat.Total), 1024) + "\n"
		html = html + "Used disk space: " + FormatFileSize(float64(diskStat.Used), 1024) + "\n"
		html = html + "Free disk space: " + FormatFileSize(float64(diskStat.Free), 1024) + "\n"
		html = html + "Persentase penggunaan space harddisk: <b>" + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%</b>\n"
		html = html + "------------------------------------------\n"
		html = html + "<b>Harap Kosongkan Ruang</b>\n"
		html = html + "1. Hapus File Backup\n"
		html = html + "2. Hapus File Log yang tidak diperlukan\n"
		html = html + "------------------------------------------\n"

		param.Set("text", html)
		param.Set("parse_mode", "HTML")
		diskObj.Telegram.SendMessage(param)
	}

	return

}

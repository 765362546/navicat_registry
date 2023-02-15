package main

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {
	fmt.Println("清理Navicat注册表信息，重置试用时间")
	fmt.Println("")
	time.Sleep(1 * time.Second)
	//删除 HKEY_CURRENT_USER\Software\PremiumSoft\Data
	data_k, _ := registry.OpenKey(registry.CURRENT_USER, `Software\PremiumSoft`, registry.ALL_ACCESS)
	defer data_k.Close()
	registry.DeleteKey(data_k, "Data")
	fmt.Println(`清理 HKEY_CURRENT_USER\Software\PremiumSoft\Data 完成`)

	// 删除 HKEY_CURRENT_USER\Software\Classes\CLSID 子项下 的info子项
	cls_k, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\CLSID`, registry.ALL_ACCESS)
	defer cls_k.Close()
	fmt.Println(`遍历 HKEY_CURRENT_USER\Software\Classes\CLSID 子项...`)
	keys, _ := cls_k.ReadSubKeyNames(50)
	for _, sub1_k := range keys {
		// fmt.Println(sub1_k)
		//下级子项
		sub_ks, _ := registry.OpenKey(cls_k, sub1_k, registry.ALL_ACCESS)
		defer sub_ks.Close()
		registry.DeleteKey(sub_ks, "Info")
	}
	fmt.Println(`清理 HKEY_CURRENT_USER\Software\Classes\CLSID 子项下 的info子项 完成`)
	fmt.Println("")
	fmt.Println("退出...")
	time.Sleep(2 * time.Second)

}

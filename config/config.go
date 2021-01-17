package config

import (
    "github.com/spf13/viper"
    "strconv"
)

// Init 创建配置文件
func Init(address, name string)error{
    viper.SetConfigFile(address)
    viper.Set("name",name)
    viper.Set("count",0)
    err := viper.WriteConfig()
    if err != nil{
        return err
    }
    return nil
}

// Read 读取配置文件
func Read(address string) ([]map[string]string, error){
    viper.SetConfigFile(address)
    err := viper.ReadInConfig()
    if err != nil{
        return nil, err
    }
    count := viper.GetInt("count")
    content:= make([]map[string]string,count)
    for i:=1;i<=count;i++{
        content[i-1] = viper.GetStringMapString(strconv.Itoa(i))
    }
    return content, nil
}

// Write 写入文件
func Write(address string, message map[string]string)error{
    viper.SetConfigFile(address)
    err := viper.ReadInConfig()
    if err != nil{
        return err
    }
    count := viper.GetInt("count")
    viper.Set("count",count+1)
    viper.Set(strconv.Itoa(count+1),message)
    err = viper.WriteConfig()
    if err != nil{
        return err
    }
    return nil
}
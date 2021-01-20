package config

import (
    "encoding/json"
    "github.com/spf13/viper"
    "reflect"
    "strconv"
)

// Init 创建配置文件
func Init(address, name string)error{
    // 打开文件
    viper.SetConfigFile(address)
    // 新建表头
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
func Write(address string, message interface{})error{
    // 打开文件
    viper.SetConfigFile(address)
    err := viper.ReadInConfig()
    if err != nil{
        return err
    }
    // 判断传入消息类型
    var mess interface{}
    t := reflect.TypeOf(message).Kind()
    switch t {
    case reflect.Map:
        mess = message.(map[string]string)
    case reflect.Slice:
        _ = json.Unmarshal(message.([]byte),&mess)
    }
    // 更新消息
    count := viper.GetInt("count")
    viper.Set("count",count+1)
    viper.Set(strconv.Itoa(count+1),mess)
    err = viper.WriteConfig()
    if err != nil{
        return err
    }
    return nil
}
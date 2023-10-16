package global

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"ry_go/src/global/conf"
	"ry_go/src/pojo"
	"time"
)

/**
 * @ClassName index
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:36
 * @Version 1.0
 */
var (
	v               *viper.Viper
	err             error
	minute          = time.Minute
	ctx             = context.Background()
	Captcha         string //redis缓存验证码前缀
	Port            string //程序使用端口
	HttpVersion     bool   //版本控制
	InterceptPrefix string
	CaptchaExp      time.Duration
	RedisClient     *redis.Client
	MysqlDClient    *gorm.DB
	EtcdClient      *clientv3.Client
	JWTKEY          string
	IpAccess        = []string{"127.0.0.1"}
	WriteList       = []string{"/api/upload/file", "/api/captcha"}
	EtcdArry        = []string{"192.168.245.22:2379"}
	FilePath        string        //静态文件上传路径
	VideoPath       string        //视频上传路径
	AdminExp        time.Duration //管理员登陆时长
	UserExp         time.Duration //用户登录时长
	CabinModel      = "auth_model.conf"
)

var (
	MysqlConfig    conf.MysqlConf    //连接mysql实例化参数
	RedisConfig    conf.RedisConf    //连接redis实例化参数
	RabbitMQConfig conf.RabbitmqConf //连接rabbitmq实例化参数
	LogConf        conf.LogCof       //连接日志实例化参数
	CabinConfig    conf.CabinConf    //连接casbin实例化参数
	ClickConfig    conf.ClickConf    //连接ck实例化参数

)

func init() {
	log.Println("实例化配置文件")
	// 构建 Viper 实例
	v = viper.New()
	v.SetConfigFile("conf.yaml") // 指定配置文件路径
	v.SetConfigName("conf")      // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")      // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	v.AddConfigPath(".") // 还可以在工作目录中查找配置
	// 查找并读取配置文件
	if err = v.ReadInConfig(); err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig() //开启监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})
	txtCon, _ := ioutil.ReadFile("banner.txt")
	fmt.Println(string(txtCon))

	viperLoadConf()

	if err = os.MkdirAll(FilePath, os.ModePerm); err != nil {
		fmt.Println("无法创建上传文件目录:", err)

	}
	if err = os.MkdirAll(VideoPath, os.ModePerm); err != nil {
		fmt.Println("无法创建上传视频目录:", err)

	}
	log.Println("所有配置完成检索，开始执行程序")

}
func viperLoadConf() {
	//读取单条配置文件
	Port = v.GetString("port")
	//设置http1.0还是2.0
	HttpVersion = v.GetBool("protocol")
	Captcha = v.GetString("captcha")
	InterceptPrefix = v.GetString("InterceptPrefix")
	CaptchaExp = time.Duration(v.GetInt("CaptchaExp")) * minute
	//日志路径及名称设置
	logConfig := v.GetStringMap("log")
	FilePath = v.GetString("FilePath")
	VideoPath = v.GetString("VideoPath")
	JWTKEY = v.GetString("JWTKEY")
	//登陆时长获取
	AdminExp = time.Duration(v.GetInt("adminExp"))
	UserExp = time.Duration(v.GetInt("userExp"))
	//读取mysql,redis,rabbitmq,casbin
	mysql := v.GetStringMap("mysql") //读取MySQL配置
	red := v.GetStringMap("redis")   //读取redis配置
	mq := v.GetStringMap("rabbitmq") //读取rabbitmq配置
	cn := v.GetStringMap("cabin")    //读取casbin配置
	ck := v.GetStringMap("click")    //读取click house配置
	//map转struct
	mapstructure.Decode(mysql, &MysqlConfig)
	mapstructure.Decode(red, &RedisConfig)
	mapstructure.Decode(mq, &RabbitMQConfig)
	mapstructure.Decode(logConfig, &LogConf)
	mapstructure.Decode(cn, &CabinConfig)
	mapstructure.Decode(ck, &ClickConfig)

	//log.Print(CabinConfig, "参数")
	//mapstructure.Decode(ca, &CaptchaConf)
	//etcd := v.GetStringSlice("etcd")
	//kafka := v.GetStringSlice("kafka")
	//oracle := v.GetStringSlice("oracle")
	//EtcdArry = append(EtcdArry, etcd...)
	//KafkaArry = append(KafkaArry, kafka...)
	log.Println("全局配置文件信息读取无误,开始载入")
	Loginit()                         //日志初始化
	Dbinit()                          //mysql初始化
	pojo.Repositoryinit(MysqlDClient) //表结构迁移
	Redisinit()                       //redis初始化

	//CabinInit() //rbac初始化
	//OracleInit()     //Oracle初始化
	//ClickhouseInit() //click house初始化
	//EtcdInit()
	//Mqinit()
}

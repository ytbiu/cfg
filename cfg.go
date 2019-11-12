package cfg

import (
	g "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"log"
)

func FromConsul(setOpts ...SetOpt) error {
	opt := &Option{}
	for _, setOpt := range setOpts {
		setOpt(opt)
	}

	if err := opt.check();err != nil{
		return err
	}

	consulSource := consul.NewSource(
		consul.WithAddress(opt.Addr),
		consul.WithPrefix(opt.ConfigPath),
		source.WithEncoder(opt.EncodeType.NewEncoder()),
	)

	if err := g.Load(consulSource); err != nil {
		return err
	}

	result := opt.Result
	if err := g.Scan(result); err != nil {
		return err
	}

	w, err := g.Watch()
	if err != nil {
		return err
	}

	go func() {
		for {
			changed, err := w.Next()
			if err != nil {
				log.Println(err)
			}

			if err := changed.Scan(result); err != nil {
				log.Println("reload config from consul err : ", err)
				continue
			}

			for _, hook := range opt.hooks {
				hook()
			}
		}
	}()

	return nil
}

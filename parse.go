package main

import (
        "launchpad.net/gozk"
        "log"
)

func zkGetNode(zkconn *zookeeper.Conn, path string, kv map[string]string) error {
        children, stat, err := zkconn.Children(path)
        if err != nil {
                log.Printf("zkGetNode:Error:  Failed getting children %s  %s", path, err.Error())
                return nil
        }
        if stat.NumChildren() == 0 {
                data, _, err := zkconn.Get(path)
                if err != nil {
                        log.Printf("zkGetNode:Error:  Failed getting data %s  %s", path, err.Error())
                        return nil
                }
                kv[path] = data
                return nil
        }
        for _,element := range children {
                zkGetNode(zkconn, path + "/" +  element,kv)
        }
        return nil
}


func main() {
  zkconn, session, err := zookeeper.Dial("127.0.0.1:2181",5e9)
  if err != nil {
    log.Printf("zkConnect:Error: Can't connect to zookeeper: %s", err.Error())
  } else {
    event := <-session
    if event.State != zookeeper.STATE_CONNECTED {
      log.Printf("zkConnect:Error: Can't reach connected state: %v", event)
      err = zkconn.Close()
      if err != nil {
        log.Printf("zkConnect:Error:  Failed closing connection to zookeeper: %s", err.Error())
      }
    } else {
      var path string = "/root" 
      var kv map[string]string
      kv = make(map[string]string)
      zkGetNode(zkconn , path , kv)
      for key, value := range kv {
        log.Printf(" %s -> %s",key, value)
      }
   }
 }
}

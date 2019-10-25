package com.github.hfantin.eurekasrv

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer

@SpringBootApplication
@EnableEurekaServer
class Application

fun main(args: Array<String>) {
	runApplication<Application>(*args)
}

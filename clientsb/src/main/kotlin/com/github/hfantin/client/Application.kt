package com.github.hfantin.client

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.cloud.client.discovery.EnableDiscoveryClient

@SpringBootApplication
@EnableDiscoveryClient
class Application

fun main(args: Array<String>) {
    runApplication<Application>(*args)
}

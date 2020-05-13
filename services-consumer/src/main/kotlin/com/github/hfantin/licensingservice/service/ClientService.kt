package com.github.hfantin.licensingservice.service

import com.github.hfantin.licensingservice.clients.Clients
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Service

@Service
class ClientService {

    @Autowired
    private lateinit var clients: Clients

    private val LOG = LoggerFactory.getLogger(javaClass.simpleName)

    @Scheduled(fixedDelay = 1000)
    fun schedulledTask() {
        val client = clients.getClient(1)
        LOG.info("clients -${client}")
    }

}
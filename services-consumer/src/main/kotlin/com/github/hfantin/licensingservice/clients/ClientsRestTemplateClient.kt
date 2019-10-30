package com.github.hfantin.licensingservice.clients

import com.github.hfantin.licensingservice.model.Client
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpMethod
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Component
import org.springframework.web.client.RestTemplate
import org.bouncycastle.crypto.tls.ConnectionEnd.client



@Component
class ClientsRestTemplateClient {

    private val LOG = LoggerFactory.getLogger(javaClass.simpleName)

    @Autowired
    private lateinit var restTemplate: RestTemplate

    private val uri = "http://clients/v1/clients"

    @Scheduled(fixedDelay = 5000)
    fun getClientsFromServiceDiscovery() {
        LOG.info("getClientsFromServiceDiscovery()")
        val clients = restTemplate.getForEntity(uri, Array<Client>::class.java)
        LOG.info("getClientsFromServiceDiscovery() clients - restExchange.getBody()=${clients}")
    }

}
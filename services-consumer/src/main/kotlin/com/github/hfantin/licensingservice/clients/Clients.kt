package com.github.hfantin.licensingservice.clients

import com.github.hfantin.licensingservice.model.Client
import org.springframework.cloud.openfeign.FeignClient
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable

@FeignClient("clients")
interface Clients {

    @GetMapping("/v1/clients")
    fun getClients(): List<Client>

    @GetMapping("/v1/clients/{id}")
    fun getClient(@PathVariable id: Int): Client?
}
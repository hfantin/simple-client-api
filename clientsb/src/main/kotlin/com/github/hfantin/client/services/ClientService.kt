package com.github.hfantin.client.services


import com.github.hfantin.client.model.ClientDto
import com.github.hfantin.client.repository.ClientRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Component

@Component
class ClientService {


    @Autowired
    private lateinit var clientRepository: ClientRepository

    fun findAll() = clientRepository.findAll()

}
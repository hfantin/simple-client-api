package com.github.hfantin.client.repository


import com.github.hfantin.client.entities.Client
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository


@Repository
interface ClientRepository : JpaRepository<Client, Int>
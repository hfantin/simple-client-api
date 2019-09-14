package com.github.hfantin.client.entities

import java.time.LocalDate
import javax.persistence.*

@Entity
@Table(name = "CLI")
data class Client(
        @Id
        @GeneratedValue(strategy = GenerationType.IDENTITY)
        var id: Int = 0,
        var name: String = "",
        @Column(name = "BIRTH_DATE")
        var birthDate: LocalDate = LocalDate.now(),
        var email: String? = null
)
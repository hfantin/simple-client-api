package com.github.hfantin.licensingservice.model


import java.time.LocalDate

data class Client(
        var id: Int = 0,
        var name: String = "",
        var birthDate: LocalDate = LocalDate.now(),
        var email: String? = null
)
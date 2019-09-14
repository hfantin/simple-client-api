package com.github.hfantin.client.model

import com.fasterxml.jackson.annotation.JsonFormat
import java.io.Serializable
import java.time.LocalDate

data class ClientDto(
        var name: String = "",
        @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "dd-MM-yyyy")
        var birthDate: LocalDate = LocalDate.now(),
        var email: String? = null
) : Serializable
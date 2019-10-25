package com.github.hfantin.zuulsrv.actuators

import org.springframework.boot.actuate.health.AbstractHealthIndicator
import org.springframework.boot.actuate.health.Health
import org.springframework.stereotype.Component

@Component
class MyHealthActuator : AbstractHealthIndicator() {
    override fun doHealthCheck(builder: Health.Builder?) {
        builder
                ?.up()
                ?.withDetail("app", "I am Alive")
                ?.withDetail("error", "Nothing!")
    }
}
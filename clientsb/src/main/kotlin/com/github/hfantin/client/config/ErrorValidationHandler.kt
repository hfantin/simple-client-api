package com.github.hfantin.client.config

import com.github.hfantin.client.model.ErrorFormDto
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.MessageSource
import org.springframework.context.i18n.LocaleContextHolder
import org.springframework.http.HttpStatus
import org.springframework.web.bind.MethodArgumentNotValidException
import org.springframework.web.bind.MissingServletRequestParameterException
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.bind.annotation.ResponseStatus
import org.springframework.web.bind.annotation.RestControllerAdvice
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException

@RestControllerAdvice
class ErrorValidationHandler {


    @Autowired
    private lateinit var messageSource: MessageSource

    @ResponseStatus(code = HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentNotValidException::class)
    fun handle(exception: MethodArgumentNotValidException) =
            exception.bindingResult.fieldErrors.map { ErrorFormDto(it.field, messageSource.getMessage(it, LocaleContextHolder.getLocale())) }


    @ResponseStatus(code = HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentTypeMismatchException::class)
    fun handle(exception: MethodArgumentTypeMismatchException) = ErrorFormDto(exception.name, exception.message)


    @ResponseStatus(code = HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MissingServletRequestParameterException::class)
    fun handle(exception: MissingServletRequestParameterException) = ErrorFormDto(exception.parameterName, exception.message)

}
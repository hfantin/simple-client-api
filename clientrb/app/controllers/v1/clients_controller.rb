module V1
    class ClientsController < ApplicationController   
        # Listar todos os clientes
        def index
            clients = Client.order('id ASC');
            render json: clients ,status: :ok
        end
    end
end
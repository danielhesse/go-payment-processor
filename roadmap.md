# Caso de Uso: "Process Transaction"

[x] Receberá os dados de uma transação
[x] Criará uma transação
  - Se o cartão de crédito for inválido:
    [] Os dados da transação serão inseridos no banco de dados com o    status="rejected" contendo a mensagem do erro
    [] A transação será publicada no Kafka

  - Caso a transação não seja aprovada:
    [] Os dados da transação serão inseridos no banco de dados com o status="rejected" contendo a mensagem do erro
    [] A transação será publicada no Kafka

  - Caso a transação seja aprovada:
    [] Os dados da transação serão inseridos no banco de dados com o status="approved"
    [] A transação será publicada no Kafka
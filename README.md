# EDA - Event Driven Architecture
## O que é?
"Arquitetura baseada a Eventos é um padrão arquitetural que promove a produção, detecção, consumo e reação a eventos. Em uma arquitetura orientada a eventos, os componentes do sistema se comunicam entre si através da emissão e escuta de eventos, permitindo uma maior desacoplagem e escalabilidade."

## Características Principais de Eventos
- Situações que ocorrem no **passado**.
- Normalmente deixa efeitos colaterais. 
    - **Exemplo:** Porta do carro abriu. Ligar a luz interna.
- Podem ser **simples** ou **complexos**.
    - **Exemplo Simples:** Usuário fez login.
    - **Exemplo Complexo:** Usuário fez uma compra.
- Podem trabalhar e forma internalizada no software ou externalizada.
    - **Exemplo Internalizado:** Evento de sistema (ex: Listener).
    - **Exemplo Externalizado:** Evento de negócio (ex: Mensageria, Streaming, etc).
- **Domain Events:** Eventos de domínio que representam algo que aconteceu no negócio. Em outras palavras, são mudanças que ocorreram no estado interno da aplicação / regra de negócio -> Ex: agregados.
    - **Exemplo:** PedidoRealizado, PagamentoAprovado, UsuárioRegistrado.

## Tipos Comuns de Eventos
- **Event Notification (Notificação de Evento):** Indica que algo aconteceu, mas não requer uma ação específica. É uma forma curta de comunicação entre componentes.
    - **Exemplo:** Um sistema de monitoramento que envia uma notificação quando um servidor fica offline.
    - **Analogia:** Alguém escreve um bilhete e entrega para você informando que a reunião foi adiada. O bilhete pode ser jogado fora após a leitura, pois não há necessidade de ação adicional. Você apenas toma conhecimento da informação e segue com suas atividades normais, podendo decidir se precisa ou não fazer algo a respeito.
- **Event-Carried State Transfer (Transferência de Estado Transportado pelo Evento):** O evento carrega informações sobre o estado que mudou, permitindo que outros componentes atualizem seu próprio estado com base nessas informações.
    - **Exemplo:** Um sistema de inventário que atualiza a quantidade de produtos com base em eventos de vendas.
    - **Analogia:** Alguém escreve um bilhete informando que a reunião foi adiada e inclui a nova data e hora. Você lê o bilhete e atualiza sua agenda com as novas informações, garantindo que você esteja ciente da mudança e possa se preparar adequadamente para a reunião na nova data.
    - **Streaming de Dados:** Tecnologias como Apache Kafka, Amazon Kinesis e Apache Pulsar são exemplos de plataformas de streaming de dados que suportam a transferência de estado transportado por eventos. Essas plataformas permitem que os eventos sejam transmitidos em tempo real, possibilitando que os sistemas consumidores atualizem seu estado com base nas informações contidas nos eventos.
    - **Vantagens:** Desacoplamento entre produtores e consumidores, escalabilidade, resiliência, flexibilidade, processamento assíncrono.
    - **Desvantagens:** Complexidade adicional, latência, garantia de entrega, monitoramento e depuração.
- **Event Sourcing (Fonte de Eventos):** Em vez de armazenar o estado atual de um sistema, todos os eventos que levaram ao estado atual são armazenados. O estado pode ser reconstruído reproduzindo esses eventos.
    - **Exemplo:** Um sistema bancário que mantém um registro de todas as transações para reconstruir o saldo da conta.
    - **Analogia:** Imagine que você está mantendo um diário detalhado de todas as suas atividades diárias. Em vez de apenas anotar o saldo final do seu dinheiro no final do mês, você registra cada transação que fez, como compras, depósitos e saques. Se você quiser saber quanto dinheiro tinha em um determinado dia, você pode revisar seu diário e somar todas as transações até aquele ponto para calcular o saldo exato.
    - **Vantagens:** Histórico completo, auditabilidade, facilidade de recuperação de estado, suporte a CQRS.
    - **Desvantagens:** Complexidade na implementação, armazenamento aumentado, desafios na modelagem de eventos.
- **Event Collaboration (Colaboração de Eventos):** Vários serviços ou componentes trabalham juntos para processar eventos e realizar ações complexas. Nessa abordagem, os serviços partem do princípio de que eles já possuem todas as informações necessárias para processar o evento e tomar decisões com base nele.
    - **Exemplo:** Um sistema de comércio eletrônico onde o processamento de um pedido envolve vários serviços, como inventário, pagamento e envio.
    - **Analogia:** Imagine que você está organizando uma festa surpresa para um amigo. Você coordena com vários amigos para garantir que todos saibam o que fazer: um amigo cuida da decoração, outro prepara a comida, outro convida os convidados, e assim por diante. Cada amigo tem uma tarefa específica, mas todos trabalham juntos para garantir que a festa seja um sucesso. Quando chega o momento da festa, todos executam suas tarefas de forma coordenada para criar uma experiência memorável para o aniversariante.
    - **Vantagens:** Desacoplamento, escalabilidade, flexibilidade, resiliência.
    - **Desvantagens:** Complexidade na coordenação, latência, monitoramento e depuração. Se um serviço demorar um pouco mais para procesar, pode impactar o outro serviço que necessitava daquela informação para continuar o processamento de forma consistente.

## Vamos falar sobre CQRS?
"CQRS (Command Query Responsibility Segregation) é um padrão arquitetural que separa as operações de leitura (queries) das operações de escrita (commands) em um sistema. Em vez de usar o mesmo modelo para ler e escrever dados, o CQRS utiliza modelos distintos para cada tipo de operação. No comando não temos retorno, ele apenas muda o estado do sistema. Na query, apenas retornamos dados, sem alterar o estado do sistema."

### CQRS vs CQS
A principal diferença é o nível de granularidade em que a separação ocorre:
- **CQS (Command Query Separation):** Princípio que sugere que um método deve ser ou um comando (que altera o estado) ou uma consulta (que retorna dados), mas não ambos.
- **CQRS:** Extensão do CQS que aplica a separação em nível arquitetural, utilizando modelos diferentes para leitura e escrita.

### CQRS e Separação Física de Dados
No CQRS, a separação entre comandos e consultas pode ser implementada de duas maneiras:
- **Separação Lógica:** Utiliza o mesmo banco de dados, mas com modelos diferentes para leitura e escrita.
- **Separação Física:** Utiliza bancos de dados diferentes para leitura e escrita, otimizando cada um para suas respectivas operações.
### CQRS e Event Sourcing
O CQRS é frequentemente combinado com o Event Sourcing, onde os comandos geram eventos que são armazenados e usados para reconstruir o estado do sistema. Isso permite uma maior auditabilidade e flexibilidade na gestão do estado. No entanto, é importante notar que o CQRS pode ser implementado sem Event Sourcing, dependendo das necessidades do sistema.
### Event Sourcing vs Command Sourcing
- **Event Sourcing:** Armazena todos os eventos que ocorreram no sistema para reconstruir o estado atual.
- **Command Sourcing:** Armazena os comandos que foram enviados ao sistema, permitindo a reprodução das ações que levaram ao estado atual.
### Como implementar CQRS?
1. **Definir Modelos Separados:** Crie modelos distintos para leitura e escrita, garantindo que cada um seja otimizado para suas operações específicas.
2. **Implementar Handlers:** Desenvolva handlers separados para comandos e consultas, garantindo que cada um lide com suas responsabilidades de forma isolada.
3. **Sincronização de Dados:** Estabeleça mecanismos para sincronizar os dados entre os modelos de leitura e escrita, garantindo consistência.
4. **Monitoramento e Logging:** Implemente monitoramento e logging para rastrear o desempenho e identificar possíveis problemas.
5. **Testes Rigorosos:** Realize testes rigorosos para garantir que ambos os modelos funcionem corretamente e que a sincronização de dados seja eficaz.
### Exemplo de implementação simples de CQRS
```csharp
public class CreateOrderCommand
{
    public Guid OrderId { get; set; }
    public string ProductName { get; set; }
    public int Quantity { get; set; }
}
public class OrderQuery
{
    public Guid OrderId { get; set; }
}
public class OrderCommandHandler
{
    public void Handle(CreateOrderCommand command)
    {
        // Lógica para criar um pedido
        Console.WriteLine($"Pedido {command.OrderId} criado para {command.ProductName} (Quantidade: {command.Quantity})");
    }
}
public class OrderQueryHandler
{
    public Order GetOrder(OrderQuery query)
    {
        // Lógica para recuperar um pedido
        return new Order { OrderId = query.OrderId, ProductName = "Exemplo", Quantity = 1 };
    }
}
public class Order
{
    public Guid OrderId { get; set; }
    public string ProductName { get; set; }
    public int Quantity { get; set; }
}
```
### Exemplos de Tecnologias para CQRS
- **MediatR:** Biblioteca para .NET que facilita a implementação do padrão CQRS
- **Axon Framework:** Framework para Java que suporta CQRS e Event Sourcing
- **EventStore:** Banco de dados especializado em Event Sourcing que pode ser usado em conjunto com CQRS
- **Apache Kafka:** Plataforma de streaming que pode ser usada para transmitir eventos entre componentes CQRS.
## Conclusão
A Arquitetura Baseada em Eventos (EDA) é um padrão arquitetural poderoso que promove a comunicação assíncrona entre componentes de um sistema através da emissão e escuta de eventos. Compreender os diferentes tipos de eventos, como Notificação de Evento, Transferência de Estado Transportado pelo Evento, Fonte de Eventos e Colaboração de Eventos, é essencial para projetar sistemas escaláveis e resilientes. Além disso, a combinação do EDA com o padrão CQRS pode trazer benefícios adicionais, como a separação clara entre operações de leitura e escrita, melhorando a performance e a flexibilidade do sistema. Ao implementar essas abordagens, é importante considerar as vantagens e desvantagens associadas, bem como as necessidades específicas do negócio para garantir uma arquitetura eficaz e eficiente.
## Refêrencias
- https://martinfowler.com/eaaDev/EventDrivenArchitecture.html
- https://microservices.io/patterns/data/event-sourcing.html
- https://docs.microsoft.com/en-us/azure/architecture/patterns/cqrs

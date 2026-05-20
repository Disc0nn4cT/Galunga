```mermaid
graph TD
    %% Клієнтський рівень
    Client([Клієнтські додатки]) --> API_Gateway[API Gateway / Nginx]

    %% Шлюз маршрутизує запити до відповідних мікросервісів
    API_Gateway -->|/api/auth/*| Auth_Service[Auth & Users Service]
    API_Gateway -->|/api/images/*| Image_Service[Image Storage Service]
    API_Gateway -->|/api/documents/*| Document_Service[PDF Document Service]

    %% Кожен сервіс має власне незалежне сховище (правило мікросервісів)
    Auth_Service --> DB_Users[(PostgreSQL: Users DB)]
    Image_Service --> DB_Images[(PostgreSQL: Images DB)]
    Image_Service --> S3_Storage[(S3 / Local Volume)]
    
    %% Асинхронна комунікація через брокер повідомлень (замість прямого виклику в пам'яті)
    Auth_Service -.->|Подія: UserRegistered| Message_Broker[[RabbitMQ / Kafka]]
    Message_Broker -.->|Обробляє подію| Email_Service[Email Worker Service]
    
    %% Сервіс розсилки працює незалежно
    Email_Service --> SMTP_Relay[SMTP Server / MailHog]

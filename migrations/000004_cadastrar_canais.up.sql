-- Cadastra os canais em que a loja vende. Eles mudam pouco, então já nascem com a
-- API: não existe rota para criar canal. As taxas são próximas das reais, e o
-- estoque de segurança 2 faz os marketplaces nunca venderem as duas últimas peças (R3).
INSERT INTO canais (nome, taxa_percentual, estoque_seguranca) VALUES
    -- A loja física: sem comissão e sem atraso para atualizar o estoque.
    ('Balcão', 0, 0),
    -- Marketplace: fica com 16% de cada venda e atualiza o estoque com atraso.
    ('Mercado Livre', 16, 2),
    -- Marketplace: fica com 20% de cada venda e atualiza o estoque com atraso.
    ('Shopee', 20, 2);

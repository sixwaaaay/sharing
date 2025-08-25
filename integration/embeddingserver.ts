// 定义响应结构
type EmbeddingData = {
    object: "embedding";
    embedding: number[];
    index: number;
};

type EmbeddingResponse = {
    object: "list";
    data: EmbeddingData[];
    model: string;
    usage: {
        prompt_tokens: number;
        total_tokens: number;
    };
};

// 定义支持的模型及其维度
const modelDimensions: Record<string, number> = {
    "text-embedding-ada-002": 1536,
    "text-similarity-ada-001": 1024,
    "text-search-ada-001": 1024,
    "code-search-ada-code-001": 1024,
    "simple-vector": 24
};

// 生成随机嵌入向量
function generateEmbedding(dimensions: number): number[] {
    return Array.from({ length: dimensions }, () => Math.random() * 2 - 1);
}

// 简单的 token 计数器
function countTokens(text: string | string[]): number {
    if (Array.isArray(text)) {
        return text.reduce((acc, t) => acc + countTokens(t), 0);
    }
    return text.split(/\s+/).length;
}

// 创建 HTTP 服务器
const server = Bun.serve({
    port: 8000,
    fetch(req) {
        const url = new URL(req.url);

        // 处理 Embedding API 请求
        if (req.method === "POST" && url.pathname === "/v1/embeddings") {
            return handleEmbeddingRequest(req);
        }

        // 默认返回 404
        return new Response(JSON.stringify({ error: "Not Found" }), {
            status: 404,
            headers: { "Content-Type": "application/json" },
        });
    },
});

// 处理嵌入请求
async function handleEmbeddingRequest(req: Request): Promise<Response> {
    try {
        // 解析请求体
        const { input, model = "text-embedding-ada-002" } = await req.json();

        // 验证输入
        if (input === undefined) {
            return new Response(JSON.stringify({ error: "Missing 'input' parameter" }), {
                status: 400,
                headers: { "Content-Type": "application/json" },
            });
        }

        // 确保 input 是数组
        const inputList = Array.isArray(input) ? input : [input];

        // 获取模型维度
        const dimensions = modelDimensions[model] || 1536;

        // 生成嵌入向量
        const data = inputList.map((text, index) => ({
            object: "embedding",
            embedding: generateEmbedding(dimensions),
            index,
        })) as EmbeddingData[];

        // 计算 token 使用量
        const promptTokens = countTokens(inputList);

        // 构建响应
        const response: EmbeddingResponse = {
            object: "list",
            data,
            model,
            usage: {
                prompt_tokens: promptTokens,
                total_tokens: promptTokens,
            },
        };

        return new Response(JSON.stringify(response), {
            headers: { "Content-Type": "application/json" },
        });
    } catch (error) {
        console.error("Error handling request:", error);
        return new Response(JSON.stringify({ error: "Internal Server Error" }), {
            status: 500,
            headers: { "Content-Type": "application/json" },
        });
    }
}

console.log(`Server running at http://localhost:${server.port}`);  
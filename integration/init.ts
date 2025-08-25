/* 
    the following code is used to initialize meilisearch search engine
*/
const meilisearchUrl = "http://localhost:7700"
const meilisearchApiKey = "masterKey|masterKey|masterKey|masterKey"

// 创建 videos 的index
const videosIndex = await fetch(`${meilisearchUrl}/indexes`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${meilisearchApiKey}`,
    },
    body: JSON.stringify({
        uid: 'videos',
        primaryKey: 'id',
    }),
})

if (!videosIndex.ok) {
    throw new Error(`创建 videos index 失败`)
}


// 配置 embedder
const embedder = `{
  "default": {
    "source": "rest",
    "url": "http://bunserver:8000/v1/embeddings",
    "request": {
      "model": "simple-vector",
      "input": ["{{text}}"]
    },
    "response": {
      "data": [
        {
            "embedding": "{{embedding}}"
        }
      ]
    },
    "documentTemplate": "{{doc.title}}"
  }
}`

const embeddersResponse = await fetch(`${meilisearchUrl}/indexes/videos/settings/embedders`, {
    method: 'PATCH',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${meilisearchApiKey}`,
    },
    body: embedder,
})

if (!embeddersResponse.ok) {
    throw new Error(`配置 embedders 失败`)
}


// 插入 数据
function* generator(start: number, count: number) {
    for (let i = start; i < start + count; i++) {
        yield { id: i, title: `title_${i}` }
    }
}

const insertResponse = await fetch(`${meilisearchUrl}/indexes/videos/documents`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${meilisearchApiKey}`,
    },
    body: JSON.stringify(Array.from(generator(1, 100)))
})

if (!insertResponse.ok) {
    throw new Error(`插入数据失败`)
}




/* 

the following is used to initialize qdrant vector database

*/

const qdrantUrl = "http://localhost:6333"
const apikey = 'qdrant|qdrant|qdrant|qdrant|qdrant|qdrant'


/* 
https://api.qdrant.tech/master/api-reference/collections/create-collection
*/
const collectionName = 'videos'

const createResp = await fetch(`${qdrantUrl}/collections/${collectionName}`, {
    method: 'PUT',
    headers: {
        'api-key': apikey,
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        vectors: {
            size: 1024,
            distance: 'Cosine',
        }
    }),
})

if (!createResp.ok) {
    throw new Error(`创建 collection 失败`)
}


function generateEmbedding(dimensions: number): number[] {
    return Array.from({ length: dimensions }, () => Math.random() * 2 - 1);
}

type Point = {
    id: number
    payload: {
        id: number;
        title: string;
    }
    vector: number[]
}

function generatePoint(id: number): Point {
    return {
        id: id,
        payload: {
            id: id,
            title: `title_${id}`,
        },
        vector: generateEmbedding(1024),
    }
}


// id from 1 to 100
const points = Array.from({ length: 100 }, (_, i) => generatePoint(i + 1))

/* 
reference: https://api.qdrant.tech/master/api-reference/points/upsert-points
*/
const resp = await fetch(`${qdrantUrl}/collections/${collectionName}/points`, {
    method: 'PUT',
    headers: {
        'api-key': apikey,
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        points,
    }),
})

if (!resp.ok) {
    throw new Error(`插入 points 失败`)
}
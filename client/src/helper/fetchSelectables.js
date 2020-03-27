import config from './config'

//int => [{}]
export default async function fetchSelectables(categoryTypeId) {
    try {
        const r = await fetch(config.host + config.apiUrl.getSelectables);
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('fetchSelectables Error:', e);
    }

    return false;
}
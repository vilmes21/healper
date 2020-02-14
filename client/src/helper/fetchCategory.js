import config from './config'

//int => [{}]
export default async function fetchCategory(categoryTypeId) {
    try {
        const r = await fetch(config.host + '/category/' + categoryTypeId);
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('fetchCategory Error:', e);
    }

    return false;
}
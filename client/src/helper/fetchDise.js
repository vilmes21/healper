import config from './config'

//int => {}
export default async function fetchDise(id) {
    try {
        const r = await fetch(`${config.host}${config.apiUrl.getDise}${id}`);
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('fetchDise Error:', e);
    }

    return false;
}
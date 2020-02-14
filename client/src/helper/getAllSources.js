import config from './config'

export default async function getAllSources() {
    try {
        const r = await fetch(config.host + '/source/index');
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('fetch Error:', e);
    }

    return false;
}
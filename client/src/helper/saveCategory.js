import config from './config'

//obj
export default async function saveCategory(din) {
    try {
        const r = await fetch(config.host + config.apiUrl.createCategory, {
            method: 'POST',
            body: JSON.stringify(din)
        });
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('saveCategory Error:', e);
    }

    return false;
}

class MsaApi {
    static async getAll() {
        const r = new Request('http://localhost:1337/list');
        return fetch(r).then(response => response.json());
    }

    static createMsa(msa) {
        const options = {
            method: 'POST',
            body: JSON.stringify(msa),
        };
        const r = new Request('http://localhost:1337/msa', options);
        return fetch(r)
            .then(async res => {
                if (res.status < 300 && res.status >= 200) return await res.json();
                throw new Error(await res.text());
            })
    }
}
export default MsaApi;

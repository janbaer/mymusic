import { buildMP3DBApiUrl } from '~/app/helpers/api-url-builder.helper';

class SongsService {
  constructor() {
    this.searchApiUrl = buildMP3DBApiUrl();
  }

  async search(searchTerm, searchField) {
    let query = `songs?q=${searchTerm}`;
    if (searchField) {
      query += `&${searchField}`;
    }

    const response = await fetch(`${this.searchApiUrl}/${query}`);
    return response.json();
  }

  async findDuplicates() {
    const response = await fetch(`${this.searchApiUrl}/songs/duplicates`);
    return response.json();
  }

  async delete(songId) {
    await fetch(`${this.searchApiUrl}/songs/${songId}`, { method: 'DELETE' });
  }

  async update(song) {
    const options = {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
      body: JSON.stringify(song)
    };

    await fetch(`${this.searchApiUrl}/songs/${song.id}`, options);
  }
}

export default new SongsService();

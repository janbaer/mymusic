let searchApiUrl = `${window.location.protocol}//${window.location.hostname}:8082`;
if (window.location.hostname !== 'localhost') {
  searchApiUrl = `${window.location.protocol}//${window.location.hostname}:${window.location.port}/api`;
}

class SongsService {
  async search(searchTerm, searchField) {
    let query = `songs?q=${searchTerm}`;
    if (searchField) {
      query += `&${searchField}`;
    }

    const response = await fetch(`${searchApiUrl}/${query}`);
    return response.json();
  }

  async findDuplicates() {
    const response = await fetch(`${searchApiUrl}/songs/duplicates`);
    return response.json();
  }

  async delete(songId) {
    await fetch(`${searchApiUrl}/songs/${songId}`, { method: 'DELETE' });
  }

  async update(song) {
    const options = {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
      body: JSON.stringify(song)
    };

    await fetch(`${searchApiUrl}/songs/${song.id}`, options);
  }
}

export default new SongsService();

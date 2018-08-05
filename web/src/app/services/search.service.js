let searchApiUrl = `${window.location.protocol}//${window.location.hostname}:8081`;
if (window.location.hostname !== 'localhost') {
  searchApiUrl += '/api';
}

class SearchService {
  async search(searchTerm, searchField) {
    let query = `songs?q=${searchTerm}`;
    if (searchField) {
      query += `&${searchField}`;
    }

    const response = await fetch(`${searchApiUrl}/${query}`);
    return response.json();
  }
}

export default new SearchService();

const searchApiUrl = `http://${window.location.hostname}:8082`;

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

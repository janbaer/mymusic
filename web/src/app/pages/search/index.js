import { h, Component } from 'preact';
import SearchPanel from './components/search-panel';
import SearchResult from './components/search-result';
import PageFooter from './../../components/page-footer';
import './index.less';

import songsService from './../../services/songs.service';

export default class SearchPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      songs: [],
      isBusy: false
    };
  }

  async startSearch(searchTerm, searchField) {
    this.setState({ isBusy: true }, async() => {
      const songs = await songsService.search(searchTerm, searchField);
      this.setState({ songs, isBusy: false });
    });
  }

  async findDuplicates() {
    this.setState({ isBusy: true }, async() => {
      const songs = await songsService.findDuplicates();
      this.setState({ songs, isBusy: false });
    });
  }

  async deleteSong(songId) {
    await songsService.delete(songId);

    const { songs } = this.state;

    const index = songs.findIndex(s => s.id === songId);
    songs.splice(index, 1);

    this.setState({ songs });
  }

  async changeSong(song) {
    await songsService.update(song);

    const { songs } = this.state;

    const index = songs.findIndex(s => s.id === song.id);
    songs[index] = song;

    this.setState({ songs });
  }

  render(props, { songs, isBusy }) {
    return (
      <div class="SearchPage">
        <header>
          <nav>
            <h1 class="title">MyMusic</h1>
            <SearchPanel
              onStartSearch={(searchTerm, searchField) => this.startSearch(searchTerm, searchField)}
              onFindDuplicates={() => this.findDuplicates()}
            />
          </nav>
        </header>
        <main>
          <SearchResult
            songs={songs}
            isBusy={isBusy}
            onDeleteSong={songId => this.deleteSong(songId)}
            onChangeSong={song => this.changeSong(song)}
          />
        </main>
        <footer>
          <PageFooter />
        </footer>
      </div>
    );
  }
}

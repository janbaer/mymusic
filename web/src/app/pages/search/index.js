import { h, Component } from 'preact';
import SearchPanel from './components/search-panel';
import SearchResult from './components/search-result';
import PageFooter from './../../components/page-footer';
import classes from './index.less';

import songsService from './../../services/songs.service';

export default class SearchPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      songs: []
    };
  }

  async startSearch(searchTerm, searchField) {
    const songs = await songsService.search(searchTerm, searchField);
    this.setState({ songs });
  }

  async deleteSong(songId) {
    await songsService.delete(songId);

    const { songs } = this.state;

    const index = songs.findIndex(s => s.id === songId);
    songs.splice(index, 1);

    this.setState({ songs });
  }

  render(props, { songs }) {
    return (
      <div class={classes.Page}>
        <header>
          <nav class="navbar" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
              <h1 class="title">MyMusic</h1>
            </div>
            <div class="navbar-menu">
              <SearchPanel
                onStartSearch={(searchTerm, searchField) => this.startSearch(searchTerm, searchField)}
              />
            </div>
          </nav>
        </header>
        <main>
          <SearchResult
            songs={songs}
            onDeleteSong={songId => this.deleteSong(songId)}
          />
        </main>
        <footer>
          <PageFooter />
        </footer>
      </div>
    );
  }
}

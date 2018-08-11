import { h, Component } from 'preact';
import classes from './search-result.less';
import DeleteSvg from './../../../../images/delete.svg';

export default class SearchResult extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
  }

  async deleteSong(songId) {
    if (this.props.onDeleteSong) {
      this.props.onDeleteSong(songId);
    }
  }

  renderDeleteButton(songId) {
    return (
      <button
        class="button is-white"
        onClick={() => this.deleteSong(songId)}
      >
        <DeleteSvg />
      </button>
    );
  }

  renderSong(song) {
    return (
      <tr>
        <td>
          { this.renderDeleteButton(song.id) }
        </td>
        <td>{ song.artist }</td>
        <td>{ song.album }</td>
        <td><span title={song.filePath}>{ song.title }</span></td>
      </tr>
    );
  }

  render({songs}) {
    const className = `${classes.SearchResultTable} table`;

    return (
      <table class={className}>
        <thead>
          <tr>
            <th />
            <th>Artist</th>
            <th>Album</th>
            <th>Title</th>
          </tr>
        </thead>
        <tbody>
          { songs.map(song => this.renderSong(song))}
        </tbody>
      </table>
    );
  }
}

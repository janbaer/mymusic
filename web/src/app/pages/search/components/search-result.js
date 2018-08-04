import { h, Component } from 'preact';
import classes from './search-result.less';

export default class SearchResult extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
  }

  renderSong(song) {
    return (
      <tr>
        <td>{song.artist}</td>
        <td>{song.album}</td>
        <td><span title={song.filePath}>{song.title}</span></td>
      </tr>
    );
  }

  render({songs}) {
    const className = `${classes.SearchResultTable} table`;

    return (
      <table class={className}>
        <thead>
          <tr>
            <th>Artist</th>
            <th>Album</th>
            <th>Title</th>
          </tr>
        </thead>
        <tbody>
          { songs.map(this.renderSong)}
        </tbody>
      </table>
    );
  }
}

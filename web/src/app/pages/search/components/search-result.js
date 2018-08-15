import { h, Component } from 'preact';
import './search-result.less';
import DeleteSvg from './../../../../images/delete.svg';
import EditSvg from './../../../../images/edit.svg';
import EditDialog from './edit-dialog.js';

export default class SearchResult extends Component {
  constructor(props) {
    super(props);
    this.state = {
      dialogModel: {
        show: null,
        song: null
      }
    };
  }

  async deleteSong(songId) {
    if (this.props.onDeleteSong) {
      this.props.onDeleteSong(songId);
    }
  }

  async editSong(song) {
    const { dialogModel } = this.state;
    dialogModel.song = song;
    const dialogResult = await this.state.dialogModel.show();
    if (dialogResult) {
      this.props.onChangeSong(dialogModel.song);
    }
  }

  renderDeleteButton(songId) {
    return (
      <button
        class="button is-white"
        title="Delete this song"
        onClick={() => this.deleteSong(songId)}
      >
        <DeleteSvg />
      </button>
    );
  }

  renderEditButton(song) {
    return (
      <button
        class="button is-white"
        title="Change this song"
        onClick={() => this.editSong(song)}
      >
        <EditSvg />
      </button>
    );
  }

  renderSong(song) {
    return (
      <tr>
        <td class="SearchResult-actionButtonsColumn">
          { this.renderDeleteButton(song.id) }
          { this.renderEditButton(song) }
        </td>
        <td>{ song.artist }</td>
        <td>{ song.album }</td>
        <td><span title={song.filePath}>{ song.title }</span></td>
      </tr>
    );
  }

  render({ songs }, { dialogModel }) {
    return (
      <table class="SearchResult-table table">
        <EditDialog dialogModel={dialogModel} />
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

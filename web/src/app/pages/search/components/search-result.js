import { h, Component } from 'preact';
import ReactPaginate from 'react-paginate';

import { buildMP3DBApiUrl } from '~/app/helpers/api-url-builder.helper';

import EditDialog from './edit-dialog.js';

import DeleteSvg from './../../../../images/delete.svg';
import EditSvg from './../../../../images/edit.svg';
import SyncSvg from './../../../../images/sync.svg';

import './search-result.less';

const PAGE_SIZE = 20;

export default class SearchResult extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isEditDialogShown: false,
      song: null,
      selectedPage: 1
    };

    this.onHideDialog = this.onHideDialog.bind(this);
  }

  async deleteSong(songId) {
    if (this.props.onDeleteSong) {
      this.props.onDeleteSong(songId);
    }
  }

  async editSong(song) {
    this.setState({ song, isEditDialogShown: true });
  }

  onHideDialog(dialogResult, song) {
    if (dialogResult) {
      this.props.onChangeSong(song);
    }
    this.setState({ song: null, isEditDialogShown: false });
  }

  resetPaginator() {
    return new Promise(resolve => {
      this.setState({ selectedPage: 1 }, resolve);
    });
  }

  buildUrlToStreamSong(songId) {
    return `${buildMP3DBApiUrl()}/songs/${songId}/content`;
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
          <div>
            { this.renderDeleteButton(song.id) }
            { this.renderEditButton(song) }
          </div>
        </td>
        <td>{ song.artist }</td>
        <td>
          <a title={song.filePath} href={this.buildUrlToStreamSong(song.id)} target="_blank">{ song.title }</a>
        </td>
        <td>{ song.album }</td>
        <td>{ song.length }</td>
      </tr>
    );
  }

  handlePageClick(e) {
    this.setState({ selectedPage: e.selected + 1 });
  }

  renderPaginator(countOfSongs, pageSize, selectedPage) {
    if (countOfSongs < pageSize) {
      return null;
    }

    const pageCount = countOfSongs / pageSize;

    return (
      <tfoot>
        <tr>
          <td colspan="5">
            <ReactPaginate
              previousLabel="zurück"
              nextLabel="vor"
              breakLabel={<span href="">...</span>}
              breakClassName="break-me"
              pageCount={pageCount}
              marginPagesDisplayed={2}
              pageRangeDisplayed={5}
              onPageChange={(e) => this.handlePageClick(e)}
              forcePage={selectedPage - 1}
              containerClassName="SearchResult-paginator pagination"
              activeClassName="pagination-link is-current"
              pageLinkClassName="pagination-link"
              previousClassName="pagination-previous"
              nextClassName="pagination-next"
            />
          </td>
        </tr>
      </tfoot>
    );
  }

  renderProgressOverlay(isBusy) {
    if (!isBusy) {
      return null;
    }
    return (
      <div class="SearchResult-progressOverlay"><SyncSvg /></div>
    );
  }

  determinateSongsToVisualize(songs, selectedPage) {
    if (songs.length <= PAGE_SIZE) {
      return songs;
    }

    const startIndex = (selectedPage * PAGE_SIZE) - PAGE_SIZE;
    return songs.slice(startIndex, startIndex + PAGE_SIZE);
  }

  render({ songs, isBusy }, state) {
    const displayedSongs = this.determinateSongsToVisualize(songs, state.selectedPage);

    return (
      <div class="SearchResult-container">
        { this.renderProgressOverlay(isBusy) }
        {state.isEditDialogShown && <EditDialog song={state.song} onHideDialog={this.onHideDialog} />}
        <table class="SearchResult-table table">
          <thead>
            <tr>
              <th class="SearchResult-actionButtonsColumn" />
              <th>Artist</th>
              <th>Title</th>
              <th>Album</th>
              <th class="SearchResult-columnLength">Length</th>
            </tr>
          </thead>
          <tbody>
            { displayedSongs.map(song => this.renderSong(song))}
          </tbody>
          { this.renderPaginator(songs.length, PAGE_SIZE, state.selectedPage) }
        </table>
      </div>
    );
  }
}

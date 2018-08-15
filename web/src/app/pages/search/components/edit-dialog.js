import { h, Component } from 'preact';
import classnames from 'classnames';

import './edit-dialog.less';

export default class EditDialog extends Component {
  constructor(props) {
    super(props);
    this.state = {
      song: {},
      isOpened: false
    };

    this.promiseResolver = null;
    this.dialogResolve = null;
    this.showDialog = this.showDialog.bind(this);
  }

  componentDidMount() {
    this.props.dialogModel.show = this.showDialog;
  }

  showDialog() {
    const song = { ...this.props.dialogModel.song };
    this.setState({ song, isOpened: true });
    return new Promise(resolve => {
      this.promiseResolver = resolve;
    });
  }

  hide(dialogResult) {
    if (dialogResult) {
      this.props.dialogModel.song = this.state.song;
    }
    this.promiseResolver(dialogResult);
    this.setState({ song: {}, isOpened: false });
  }

  changeSongProperty(property, value) {
    const { song } = this.state;
    song[property] = value;
    this.setState({ song });
  }

  render({}, { song, isOpened }) { // eslint-disable-line no-empty-pattern
    const backgroundContainerClassname = classnames(
      'EditDialog-backgroundContainer',
      { 'is-visible': isOpened }
    );

    return (
      <div class={backgroundContainerClassname}>
        <dialog class="EditDialog-dialog" open={isOpened}>
          <label class="label" for="artist">Artist</label>
          <input class="input" name="artist" type="text" value={song.artist} onChange={e => this.changeSongProperty('artist', e.target.value)} />

          <label class="label" for="title">Title</label>
          <input class="input" name="title" type="text" value={song.title} onChange={e => this.changeSongProperty('title', e.target.value)} />

          <label class="label" for="album">Album</label>
          <input class="input "name="album" type="text" value={song.album} onChange={e => this.changeSongProperty('album', e.target.value)} />

          <label class="label" for="genre">Genre</label>
          <input class="input" name="genre" type="text" value={song.genre} onChange={e => this.changeSongProperty('genre', e.target.value)} />

          <button class="button" onClick={() => this.hide(false)}>Cancel</button>
          <button class="button is-primary" onClick={() => this.hide(true)}>Ok</button>
        </dialog>
      </div>
    );
  }
}

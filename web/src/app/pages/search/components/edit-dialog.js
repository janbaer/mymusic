import { h, Component } from 'preact';

import './edit-dialog.less';

export default class EditDialog extends Component {
  constructor(props) {
    super(props);
    const song = { ...this.props.song };
    this.state = { song };

    this.hide = this.hide.bind(this);
    this.onChangeValue = this.onChangeValue.bind(this);
  }

  hide(dialogResult) {
    this.props.onHideDialog(dialogResult, this.state.song);
  }

  onChangeValue(event) {
    const { song } = this.state;
    song[event.target.name] = event.target.value;

    this.setState({ song });
  }

  render(props, { song }) {
    return (
      <div class="EditDialog-backgroundContainer">
        <dialog class="EditDialog-dialog" open>
          <label class="label" for="artist">Artist</label>
          <input class="input" name="artist" type="text"
            value={song.artist} onChange={this.onChangeValue} />

          <label class="label" for="title">Title</label>
          <input class="input" name="title" type="text"
            value={song.title} onChange={this.onChangeValue} />

          <label class="label" for="album">Album</label>
          <input class="input" name="album" type="text"
            value={song.album} onChange={this.onChangeValue} />

          <label class="label" for="genre">Genre</label>
          <input class="input" name="genre" type="text"
            value={song.genre} onChange={this.onChangeValue} />

          <button class="button" onClick={() => this.hide(false)}>Cancel</button>
          <button class="button is-primary" onClick={() => this.hide(true)}>Ok</button>
        </dialog>
      </div>
    );
  }
}

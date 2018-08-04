import { h, Component } from 'preact';
import classes from './search-panel.less';
import SearchSvg from './../../../../images/search.svg';

export default class SearchPanel extends Component {
  constructor(props) {
    super(props);

    this.state = {
      searchField: '',
      searchTerm: '',
    };

    this.searchFields = [
      { name: '--All Fields--', value: '' },
      { name: 'Artists', value: 'artist' },
      { name: 'Title', value: 'title' },
      { name: 'Album', value: 'album' },
    ];
  }

  startSearch() {
    const { searchTerm, searchField } = this.state;
    if (searchTerm) {
      this.props.onStartSearch(searchTerm, searchField);
    }
  }

  handleSearchInputKeyPress(key, value) {
    if (key === 'Enter') {
      this.setState({searchTerm: value}, this.startSearch);
    }
  }

  renderSearchInput(searchTerm) {
    return (
      <div class="navbar-item">
        <input
          class="input"
          type="text"
          placeholder="Enter the search term"
          value={searchTerm}
          onKeyPress={event => this.handleSearchInputKeyPress(event.key, event.target.value)}
          onChange={event => this.setState({searchTerm: event.target.value})}
        />
      </div>
    );
  }

  renderSearchFields(searchField) {
    return (
      <div class="navbar-item">
        <div class="field">
          <div class="control">
            <div class="select">
              <select
                value={searchField}
                onChange={event => this.setState({ searchField: event.target.value })}
              >
                {
                  this.searchFields.map(searchField =>
                    <option value={searchField.value}>{searchField.name}</option>
                  )}
              </select>
            </div>
          </div>
        </div>
      </div>
    );
  }

  renderSearchButton() {
    return (
      <div class="navbar-item">
        <a
          class="button is-primary is-rounded"
          onClick={() => this.startSearch()}
        >
          <SearchSvg class={classes.SearchIcon} />
        </a>
      </div>
    );
  }

  render(props, state) {
    return (
      <div class="navbar-end">
        { this.renderSearchInput(state.searchTerm) }
        { this.renderSearchFields(state.searchField) }
        { this.renderSearchButton() }
      </div>
    );
  }
}

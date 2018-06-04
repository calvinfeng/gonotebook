/**
 * @author Calvin Feng
 */

// Libraries
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import axios from 'axios';
import Button from '@material-ui/core/Button';

// Helpers
import { classifyImageFile } from './util';

// Style
import './index.scss';


interface IndexState {
    currentImage: HTMLImageElement;
    resultMessage: string | undefined;
    errorMessage: string | undefined;
}

class Index extends React.Component<any, IndexState> {

    private fileReader: FileReader

    constructor(props) {
        super(props);

        this.state = {
            currentImage: undefined,
            errorMessage: undefined,
            resultMessage: undefined
        };

        this.fileReader = new FileReader();
        this.fileReader.onload = this.handleFileOnLoad;
    }

    handleFileOnLoad = (e: ProgressEvent): void => {
        const img = new Image();
        img.src = this.fileReader.result;

        this.setState({ currentImage: img });
    }

    handleImageFileSelected = (e: React.FormEvent<HTMLInputElement>): void => {
        const imageType = /image.*/;
        const file: File = e.currentTarget.files[0];

        if (file.type.match(imageType)) {
            this.fileReader.readAsDataURL(file);

            classifyImageFile(file).then((res) => {
                this.setState({ 
                    resultMessage: res,
                    errorMessage: undefined 
                });
            }).catch((err) => {
                this.setState({
                    resultMessage: undefined,
                    errorMessage: err
                });
            });
        } else {
            this.setState({ errorMessage: "File not supported!" });
        }
    }

    get imageLoader(): JSX.Element {
        let image: JSX.Element;
        if (this.state.currentImage) {
            image = <img src={this.state.currentImage.src} />;
        } else {
            image = <div />;
        }

        return (
            <section className="image-loader">
                <div className="image-container">
                    {image}
                </div>
                <form>
                    <input className="image-input"
                            accept="image/*"
                            id="hidden-image-file-input" 
                            type="file" 
                            name="imageLoader" 
                            onChange={this.handleImageFileSelected} />
                    <label htmlFor="hidden-image-file-input">
                        <Button variant="contained" component="span">Upload</Button>
                    </label>
                </form>
            </section>
        );
    }

    get error(): JSX.Element {
        if (this.state.errorMessage) {
            return <p>{this.state.errorMessage}</p>;
        }

        return <div></div>;
    }

    get result(): JSX.Element {
        if (this.state.resultMessage) {
            return <p>{this.state.resultMessage}</p>;
        }

        return <div></div>;
    }

    render() {
        return (
            <section className="index-container">
                <h1>Welcome!</h1>
                <p>Let's try to recognize an image!</p>
                <section className="image-classifier">
                    {this.imageLoader}
                    {this.result}
                    {this.error}
                </section>
            </section>
        )
    }
}


ReactDOM.render(<Index />, document.getElementById('root'));
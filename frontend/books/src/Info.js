import React from 'react';

class Information extends React.Component {
    render() {
        const Title = 'Book Title';
        const ShowTitle = true;

        if (ShowTitle) {
            return (
                <div className="Info">
                    <h1>{Title}</h1>
                    <p>Favorite Books</p>
                </div>
            );
        }
        else {
            return (
                <p>Nothing to show</p>
            );
        }
    }
}

export default Information;


// export default function Info() {
//     const Title = 'Book Title';
//     const ShowTitle = true;
//
//     if (ShowTitle) {
//         return (
//             <div className="Info">
//                 <h1>{Title}</h1>
//                 <p>Favorite Books</p>
//             </div>
//         );
//     }
//     else {
//         return (
//             <p>Nothing to show</p>
//         );
//     }
// }
//
//  Package.swift
//  vcfgresql
//
//  Created by Rodney Dyer on 5/7/19.
//  Copyright © 2019 Rodney Dyer. All rights reserved.
//

import PackageDescription

let package = Package(
    name: "Eigen",
    dependencies: [
        .package(url: "https://github.com/nicklockwood/SwiftFormat", from: "0.35.8"),
        .package(url: "https://github.com/Realm/SwiftLint", from: "0.28.1"),
        .package(url: "https://github.com/orta/Komondor", from: "1.0.0"),
    ],
    targets: [
        // This is just an arbitrary Swift file in the app, that has
        // no dependencies outside of Foundation, the dependencies section
        
    ]
)

//// The settings for the git hooks for our repo
//#if canImport(PackageConfig)
//import PackageConfig
//
//let config = PackageConfig([
//    "komondor": [
//        // When someone has run `git commit`, first run
//        // run SwiftFormat and the auto-correcter for SwiftLint
//        "pre-commit": [
//            "swift run swiftformat .",
//            "swift run swiftlint autocorrect --path Artsy/",
//            "git add .",
//        ],
//    ]
//    ])
//#endif

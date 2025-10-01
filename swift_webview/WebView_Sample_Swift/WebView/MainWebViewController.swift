//
//  MainWebViewController.swift
//  WebView_Sample_Swift
//
//  Created by Sarah Jeong on 7/5/24.
//

import UIKit
import WebKit

class MainWebViewController: UIViewController, WKScriptMessageHandler, WKUIDelegate, WKNavigationDelegate {
    
    var containerView: UIView!
    var mainWebView: WKWebView!
    
    let sharedProcessPool = WKProcessPool()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        setContainerView()
        setWebView()
    }
    
    // MARK: ContainerView
    func setContainerView() {
     
        containerView = UIView()
        containerView.backgroundColor = .white
        view.addSubview(containerView)
        
        containerView.translatesAutoresizingMaskIntoConstraints = false
        
        NSLayoutConstraint.activate([
            containerView.topAnchor.constraint(equalTo: view.topAnchor),
            containerView.leadingAnchor.constraint(equalTo: view.leadingAnchor),
            containerView.trailingAnchor.constraint(equalTo: view.trailingAnchor),
            containerView.bottomAnchor.constraint(equalTo: view.bottomAnchor)
        ])
    }
    
    func setWebView() {
        
        // Configure WKWebViewConfiguration
        let configuration = WKWebViewConfiguration()
        
        let preferences = WKPreferences()
        preferences.javaScriptCanOpenWindowsAutomatically = true
        preferences.minimumFontSize = 10.0
        
        let webpagePreferences = WKWebpagePreferences()
        webpagePreferences.allowsContentJavaScript = true
        webpagePreferences.preferredContentMode = .mobile
        
        let userContentController = WKUserContentController()
        userContentController.add(self, name: "messageHandler")
        
        configuration.processPool = sharedProcessPool
        configuration.preferences = preferences
        configuration.defaultWebpagePreferences = webpagePreferences
        configuration.userContentController = userContentController
        
        mainWebView = WKWebView(frame: .zero, configuration: configuration)
        containerView.addSubview(mainWebView)
        
        mainWebView.translatesAutoresizingMaskIntoConstraints = false
        mainWebView.uiDelegate = self
        mainWebView.navigationDelegate = self
        
        mainWebView.scrollView.showsVerticalScrollIndicator = true
        mainWebView.allowsBackForwardNavigationGestures = true
        mainWebView.allowsLinkPreview = true
        
        #if DEBUG
        if #available(iOS 16.4, *) {
            mainWebView.isInspectable = true
        }
        #endif
        
        NSLayoutConstraint.activate([
            mainWebView.topAnchor.constraint(equalTo: containerView.topAnchor),
            mainWebView.leadingAnchor.constraint(equalTo: containerView.leadingAnchor),
            mainWebView.trailingAnchor.constraint(equalTo: containerView.trailingAnchor),
            mainWebView.bottomAnchor.constraint(equalTo: containerView.bottomAnchor)
        ])
        
        // For local testing - choose one of these approaches:
        
        // Option 1: Direct localhost (iOS Simulator only)
        // if let url = URL(string: "http://localhost:3000") {
        
        // Option 2: Mac's local IP (find with `ifconfig` or `ipconfig getifaddr en0`)
        // if let url = URL(string: "http://192.168.1.100:3000") {
        
        // Option 3: Tailscale IP (more secure and works anywhere)
        // if let url = URL(string: "http://100.x.x.x:3000") {
        
        // Current: Example URL for testing
        if let url = URL(string: "http://localhost:3000") {
            let request = URLRequest(url: url)
            mainWebView.load(request)
        }
    }

    // MARK: - WKScriptMessageHandler Method
    func userContentController(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
        
        if message.name == "messageHandler" {
            
            if let messageBody = message.body as? String {
                print("JavaScript message received: \(messageBody)")
                
                switch messageBody {
                case "login":
                    print("Login")
                    
                case "logout":
                    print("Logout")
                    
                default:
                    print("Unknown action")
                }
            }
        }
    }
    
    // MARK: - WKNavigationDelegate
    func webView(_ webView: WKWebView, didStartProvisionalNavigation navigation: WKNavigation!) {
        print("Web view started loading content")
    }
    
    func webView(_ webView: WKWebView, didReceiveServerRedirectForProvisionalNavigation navigation: WKNavigation!) {
        print("Received server redirect")
    }
    
    func webView(_ webView: WKWebView, didFailProvisionalNavigation navigation: WKNavigation!, withError error: Error) {
        print("Encountered load error: \(error.localizedDescription)")
        
        // Show alert for debugging
        DispatchQueue.main.async {
            let alert = UIAlertController(title: "Load Error", 
                                        message: error.localizedDescription, 
                                        preferredStyle: .alert)
            alert.addAction(UIAlertAction(title: "OK", style: .default))
            self.present(alert, animated: true)
        }
    }
    
    func webView(_ webView: WKWebView, didFinish navigation: WKNavigation!) {
        print("Navigation completed")
    }
    
    func webView(_ webView: WKWebView, didFail navigation: WKNavigation!, withError error: Error) {
        print("Navigation failed: \(error.localizedDescription)")
    }
    
    func webViewWebContentProcessDidTerminate(_ webView: WKWebView) {
        print("Web content process terminated - reloading...")
        
        // Automatically reload when web process crashes
        if let currentURL = webView.url {
            let request = URLRequest(url: currentURL)
            webView.load(request)
        }
    }
    
    func webView(_ webView: WKWebView, decidePolicyFor navigationAction: WKNavigationAction, decisionHandler: @escaping (WKNavigationActionPolicy) -> Void) {
        print("Navigation request: \(navigationAction.request.url?.absoluteString ?? "Unknown URL")")
        
        decisionHandler(.allow)
    }
}
